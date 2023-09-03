package ws

import (
	"github.com/cploutarchou/crypto-sdk-suite/bybit/ws/client"
	"github.com/cploutarchou/crypto-sdk-suite/bybit/ws/kline"
)

type WebSocket interface {
	Kline(channel kline.ChannelType, environment kline.Environment, subChannel kline.SubChannel) *kline.TheKline
}

type webSocketImpl struct {
	kline *kline.TheKline
}

func (w webSocketImpl) Kline(channel kline.ChannelType, environment kline.Environment, subChannel kline.SubChannel) *kline.TheKline {
	w.kline = new(kline.TheKline)
	w.kline.Config.ChannelType = channel
	w.kline.Config.Environment = environment
	w.kline.Config.SubChannel = subChannel
	return w.kline

}

func NewWebSocket(key, secretKey string) (WebSocket, error) {
	wsClient, err := client.NewWSClient(key, secretKey)
	if err != nil {
		return nil, err
	}
	k := &kline.TheKline{}
	return &webSocketImpl{
		kline: k.SetClient(wsClient),
	}, nil
}
