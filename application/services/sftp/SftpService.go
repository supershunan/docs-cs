package sftp

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/crypto/ssh"
)

// SftpService 通过 SSH/SCP 上传文件到服务器
type SftpService struct{}

func NewSftpService() *SftpService {
	return &SftpService{}
}

// UploadToServer 将本地路径（文件或目录）上传到服务器（密码认证）
// host: 服务器地址，port: SSH 端口（如 "22"），username、password: 登录凭据
// localPath: 本地文件或目录路径，remotePath: 服务器目标路径
func (s *SftpService) UploadToServer(host, port, username, password, localPath, remotePath string) string {
	if host == "" || username == "" || password == "" || localPath == "" || remotePath == "" {
		return "请填写完整：服务器地址、用户名、密码、本地路径、远程路径"
	}
	config := s.sshConfigPassword(username, password)
	return s.doUpload(host, port, config, localPath, remotePath)
}

// UploadToServerWithKey 使用 SSH 密钥认证上传（适用于服务器禁用密码登录的场景）
// privateKeyPath: 私钥文件路径（如 ~/.ssh/id_rsa），keyPassphrase: 私钥密码（无则传空）
func (s *SftpService) UploadToServerWithKey(host, port, username, privateKeyPath, keyPassphrase, localPath, remotePath string) string {
	if host == "" || username == "" || privateKeyPath == "" || localPath == "" || remotePath == "" {
		return "请填写完整：服务器地址、用户名、私钥路径、本地路径、远程路径"
	}
	config, err := s.sshConfigKey(username, privateKeyPath, keyPassphrase)
	if err != nil {
		return fmt.Sprintf("加载私钥失败: %v", err)
	}
	return s.doUpload(host, port, config, localPath, remotePath)
}

func (s *SftpService) sshConfigPassword(username, password string) *ssh.ClientConfig {
	return &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
}

func (s *SftpService) sshConfigKey(username, privateKeyPath, keyPassphrase string) (*ssh.ClientConfig, error) {
	keyBytes, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("读取私钥文件: %w", err)
	}
	var signer ssh.Signer
	if keyPassphrase != "" {
		signer, err = ssh.ParsePrivateKeyWithPassphrase(keyBytes, []byte(keyPassphrase))
	} else {
		signer, err = ssh.ParsePrivateKey(keyBytes)
	}
	if err != nil {
		return nil, fmt.Errorf("解析私钥: %w", err)
	}
	return &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}, nil
}

func (s *SftpService) doUpload(host, port string, config *ssh.ClientConfig, localPath, remotePath string) string {
	if port == "" {
		port = "22"
	}
	addr := fmt.Sprintf("%s:%s", host, port)
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return fmt.Sprintf("连接服务器失败: %v", err)
	}
	defer client.Close()

	info, err := os.Stat(localPath)
	if err != nil {
		return fmt.Sprintf("读取本地路径失败: %v", err)
	}

	if info.IsDir() {
		return s.uploadDir(client, localPath, remotePath)
	}
	return s.uploadFile(client, localPath, remotePath, info)
}

func (s *SftpService) uploadFile(client *ssh.Client, localPath, remotePath string, info os.FileInfo) string {
	session, err := client.NewSession()
	if err != nil {
		return fmt.Sprintf("创建会话失败: %v", err)
	}
	defer session.Close()

	file, err := os.Open(localPath)
	if err != nil {
		return fmt.Sprintf("打开文件失败: %v", err)
	}
	defer file.Close()

	stdin, err := session.StdinPipe()
	if err != nil {
		return fmt.Sprintf("创建管道失败: %v", err)
	}

	go func() {
		defer stdin.Close()
		mode := 0644
		if info.Mode().Perm() != 0 {
			mode = int(info.Mode().Perm())
		}
		baseName := filepath.Base(localPath)
		fmt.Fprintf(stdin, "C%04o %d %s\n", mode, info.Size(), baseName)
		io.Copy(stdin, file)
		fmt.Fprint(stdin, "\x00")
	}()

	// remotePath 为单文件时的目标目录
	remoteDir := strings.TrimSuffix(remotePath, "/")
	remoteDir = strings.ReplaceAll(remoteDir, "\\", "/")
	cmd := fmt.Sprintf("scp -t %s", remoteDir)
	if err := session.Run(cmd); err != nil {
		return fmt.Sprintf("上传失败: %v", err)
	}
	return ""
}

func (s *SftpService) uploadDir(client *ssh.Client, localPath, remotePath string) string {
	session, err := client.NewSession()
	if err != nil {
		return fmt.Sprintf("创建会话失败: %v", err)
	}
	defer session.Close()

	stdin, err := session.StdinPipe()
	if err != nil {
		return fmt.Sprintf("创建管道失败: %v", err)
	}

	remoteDir := strings.TrimSuffix(remotePath, "/")
	remoteDir = strings.ReplaceAll(remoteDir, "\\", "/")

	go func() {
		defer stdin.Close()
		s.sendDirRecursive(stdin, localPath, "")
	}()

	cmd := fmt.Sprintf("scp -rt %s", remoteDir)
	if err := session.Run(cmd); err != nil {
		return fmt.Sprintf("上传目录失败: %v", err)
	}
	return ""
}

// sendDirRecursive 按 SCP 协议递归发送目录，relDir 为当前相对路径
func (s *SftpService) sendDirRecursive(stdin io.Writer, localBase, relDir string) {
	fullPath := filepath.Join(localBase, relDir)
	entries, err := os.ReadDir(fullPath)
	if err != nil {
		return
	}
	for _, entry := range entries {
		name := entry.Name()
		subRel := name
		if relDir != "" {
			subRel = filepath.Join(relDir, name)
		}
		subRel = strings.ReplaceAll(subRel, "\\", "/")
		full := filepath.Join(localBase, subRel)
		info, err := entry.Info()
		if err != nil {
			continue
		}
		if entry.IsDir() {
			fmt.Fprintf(stdin, "D%04o 0 %s\n", 0755, name)
			s.sendDirRecursive(stdin, localBase, subRel)
			fmt.Fprint(stdin, "E\n")
		} else {
			file, err := os.Open(full)
			if err != nil {
				continue
			}
			mode := 0644
			if info.Mode().Perm() != 0 {
				mode = int(info.Mode().Perm())
			}
			fmt.Fprintf(stdin, "C%04o %d %s\n", mode, info.Size(), name)
			io.Copy(stdin, file)
			file.Close()
			fmt.Fprint(stdin, "\x00")
		}
	}
}
