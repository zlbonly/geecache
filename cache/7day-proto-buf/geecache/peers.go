package geecache

import   pb "7day-proto-buf/geecache/geecachepb"

type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

type PeerGetter interface {
	Get(in *pb.Request, out *pb.Response) error
}


