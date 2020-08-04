# go-tlsreuse

以 go-reuse 延伸對 TLS 進行支援，背後機制原理與 go-reuse 和 tls 相同，用法也大同小異。

## Examples


```Go
// listen on the same port. oh yeah.
crt, err := tls.LoadX509KeyPair("key.crt", "cert.key")
if err != nil {
  panic(err)
}
tlsConfig := &tls.Config{}
tlsConfig.Certificates = []tls.Certificate{crt}
tlsConfig.Time = time.Now
tlsConfig.Rand = rand.Reader


l1, _ := tlsreuse.Listen("tcp", "127.0.0.1:1234", tlsConfig)
l2, _ := tlsreuse.Listen("tcp", "127.0.0.1:1234", tlsConfig)
```
