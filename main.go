package main

import interact "smartContract/interact"

func main() {
	//interact.Approve("0x130C66f906e03e371D25e447E1fe7d7331631D5a", int64(1000000))
	//interact.Approve("0xB2174c3Cb47bC9E417908905B7F8D65d06f4140c", int64(1000000))
	interact.TransferFrom("0xB2174c3Cb47bC9E417908905B7F8D65d06f4140c", "0x22B570d80C34a6C9A1F4342C8D7088164Ef7B3e1", int64(10000))
}
