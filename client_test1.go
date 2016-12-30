package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"encoding/binary"
	"bytes"
	"time"
	"io"
	"xingo_demo/pb"
	"os"
	"os/signal"
	"github.com/viphxin/xingo/fnet"
	"github.com/viphxin/xingo/iface"
	"math/rand"
)

type PkgData struct {
	Len   uint32
	MsgId uint32
	Data  []byte
}


type MyPtotoc struct{
	Pid int32
}

func (this *MyPtotoc)OnConnectionMade(fconn iface.Iclient){
	fmt.Println("链接建立")
	go func() {
		for{
			msg := &pb.Talk{
				Content: "哈哈哈哈哈哈哈哈哈哈sdadasdasfas 萨达萨达撒发送的发大声大声",
			}
			this.Send(fconn, 0, msg)
			this.Send(fconn, 10, nil)
			time.Sleep(2000*time.Millisecond)
		}
	}()
}

func (this *MyPtotoc)OnConnectionLost(fconn iface.Iclient){
	fmt.Println("链接丢失")
}

func (this *MyPtotoc) Unpack(headdata []byte) (head *PkgData, err error) {
	headbuf := bytes.NewReader(headdata)

	head = &PkgData{}

	// 读取Len
	if err = binary.Read(headbuf, binary.LittleEndian, &head.Len); err != nil {
		return nil, err
	}

	// 读取MsgId
	if err = binary.Read(headbuf, binary.LittleEndian, &head.MsgId); err != nil {
		return nil, err
	}

	// 封包太大
	//if head.Len > MaxPacketSize {
	//	return nil, packageTooBig
	//}

	return head, nil
}

func (this *MyPtotoc) Pack(msgId uint32, data proto.Message) (out []byte, err error) {
	outbuff := bytes.NewBuffer([]byte{})
	// 进行编码
	dataBytes := []byte{}
	if data != nil {
		dataBytes, err = proto.Marshal(data)
	}

	if err != nil {
		fmt.Println(fmt.Sprintf("marshaling error:  %s", err))
	}
	// 写Len
	if err = binary.Write(outbuff, binary.LittleEndian, uint32(len(dataBytes))); err != nil {
		return
	}
	// 写MsgId
	if err = binary.Write(outbuff, binary.LittleEndian, msgId); err != nil {
		return
	}

	//all pkg data
	if err = binary.Write(outbuff, binary.LittleEndian, dataBytes); err != nil {
		return
	}

	out = outbuff.Bytes()
	return

}

func (this *MyPtotoc)DoMsg(fconn iface.Iclient, pdata *PkgData){
	//处理消息
	fmt.Println(fmt.Sprintf("msg id :%d, data len: %d", pdata.MsgId, pdata.Len))
	if pdata.MsgId == 1 || pdata.MsgId == 10{
		bdata := &pb.BroadCast{}
		proto.Unmarshal(pdata.Data, bdata)
		println(bdata.GetContent())
	}
}

func (this *MyPtotoc)Send(fconn iface.Iclient, msgID uint32, data proto.Message){
	dd, err := this.Pack(msgID, data)
	if err == nil{
		fconn.Send(dd)
	}else{
		fmt.Println(err)
	}

}

func (this *MyPtotoc)StartReadThread(fconn iface.Iclient){
	go func() {
		for {
		//read per head data
		headdata := make([]byte, 8)

		if _, err := io.ReadFull(fconn.GetConnection(), headdata); err != nil {
			fmt.Println(err)
			this.OnConnectionLost(fconn)
			return
		}
		pkgHead, err := this.Unpack(headdata)
		if err != nil {
			this.OnConnectionLost(fconn)
			return
		}
		//data
		if pkgHead.Len > 0 {
			pkgHead.Data = make([]byte, pkgHead.Len)
			if _, err := io.ReadFull(fconn.GetConnection(), pkgHead.Data); err != nil {
				this.OnConnectionLost(fconn)
				return
			}
		}
		this.DoMsg(fconn, pkgHead)
	}
	}()
}

func (this *MyPtotoc)AddRpcRouter(router interface{}){

}

func main() {
	nets := []int{11009, 11010, 11011, 11012}
	for i := 0; i< 10000; i ++{
		client := fnet.NewTcpClient("0.0.0.0", nets[rand.Intn(len(nets))], &MyPtotoc{})
		client.Start()
		time.Sleep(100*time.Millisecond)
	}

	// close
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	fmt.Println("=======", sig)
}