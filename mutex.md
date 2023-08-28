### mutexを使って排他的制御を実行してみた
- 調べること: global変数でないcntがincとdecで共有されていない。理屈を調べる

[URL](https://go.dev/play/p/6A2gRQqKEye)


### 
- sync.Lock実行後にパニックが起きてもUnlockが実行されるために無名関数を生成し、defer文を実行するようにしてみた

[URL](https://go.dev/play/p/borclhSmaOG)
