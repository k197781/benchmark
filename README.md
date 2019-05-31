# 計測時間
user：コマンドを実行するのに使用したユーザーCPU時間  
system：コマンドを実行するのに使用したシステムCPU時間

計測はtimeコマンドで行なった

### sum
1から500000000までの足し算
```
eg.）time ruby sum.rb
```
||user|system|
|:---|:---|:---|
|go|0.40s|0.17s|
|ruby|21.90s|0.20s|
|python|48.56s|0.08s|
