package main

import "client/tcp_util"

func main() {
	tcp_util.Start("localhost", "443")
}
