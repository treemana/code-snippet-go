package main

import (
	"context"
	"os/exec"
)

// 执行命令行
func execCommand(ctx context.Context, name string, arg ...string) error {
	cmd := exec.CommandContext(ctx, name, arg...)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := execCommand(context.Background(), "ls", "-al"); err != nil {
		panic(err)
	}
}
