package executer

import (
	"os/exec"
	"fmt"
	"bufio"
	"os"
	"strings"
)

type TcpdumpExecuter struct {
	cmd *exec.Cmd
}

func (this *TcpdumpExecuter) RunTcpdump(pcapDirectory, filename, commandOptions string) error {
	fmt.Println("starting tcpdump")

	tcpdumpCommand := append(
		[]string{"shell", "tcpdump", "-w", pcapDirectory + filename},
		strings.Split(commandOptions, " ")...)

	fmt.Println("tcpdump command:", tcpdumpCommand)

	this.cmd = exec.Command("bin/adb/adb.exe", tcpdumpCommand...)

	if err := this.cmd.Start(); err != nil {
		return err
	}

	fmt.Println("tcpdump running")

	return nil
}

func (this *TcpdumpExecuter) TerminateTcpdump() error {
	fmt.Println("press 'q' to stop tcpdump")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "q" {
			fmt.Println("stopping tcpdump")

			if err := this.cmd.Process.Kill(); err != nil {
				return err
			}

			fmt.Println("tcpdump stopped")
			break
		}
	}

	return nil
}

func (this *TcpdumpExecuter) AdbPullPcapFile(pcapDirectory, filename string) error {
	fmt.Println("retrieving pcap file")

	cmd := exec.Command(
		"bin/adb/adb.exe",
		"pull",
		pcapDirectory + filename,
		"pcap/" + filename)

	if err := cmd.Run(); err != nil {
		return err
	}

	fmt.Println("pcap file retrieved")

	return nil
}

func (this *TcpdumpExecuter) OpenWithWireshark(wiresharkDirectory, filename string) error {
	fmt.Println("opening pcap file with Wireshark")

	cmd := exec.Command(
		wiresharkDirectory + "Wireshark.exe",
		"-r",
		"pcap/" + filename)

	if err := cmd.Run(); err != nil {
		return err
	}

	fmt.Println("pcap opened successfully with Wireshark")

	return nil
}