# 使い方

# オプション
  -e, --elev int     elevatorNumber option (default 5) エレベータの個数。

  -f, --floor int    floorNumber option (default 10) 建物の階数。

  -h, --help         help for elevator-simulator

  -m, --max int      elevatorMaxPeople option (default 10) エレベータに乗れる最大人数。

  -r, --rate float   createHumanRate option (default 1) 人の生成速度。数字を大きくすると増えやすくなる。

  -s, --sleep int    sleep option (default 100) 一ステップごとの待機時間（ミリ秒）。

  -t, --time int     endTime option (default 100) 最大ステップ数。


# 例
```bash
./elevator-simulator --time 1000 --floor 15 --elev 10 --rate 5.0 --max 20 --sleep 10
./elevator-simulator -t 1000 -f 15 -e 10 -r 5.0 -m 20 -s 10
```

windowsの場合はcmdから実行できます。
linuxの場合はファイルに実行権限を付与する必要があります。
```bash
chmod u+x ./elevator-simulator
```

# 参考
https://qiita.com/NorsteinBekkler/items/b2418cd5e14a52189d19
https://code.visualstudio.com/docs/remote/containers#_sharing-git-credentials-with-your-container
http://psychedelicnekopunch.com/archives/1780
https://golang.hateblo.jp/entry/2019/10/07/171630
https://qiita.com/quicksort/items/c9522793a941edf074fd
https://inabajunmr.hatenablog.com/entry/2019/11/07/093217
https://blog.junkata.com/posts/go-map-for-random
https://qiita.com/egnr-in-6matroom/items/282aa2fd117aab9469bd
https://qiita.com/crifff/items/b116e6235fedcd18e0de
https://blog.y-yuki.net/entry/2017/05/06/000000
https://developer.fyne.io/tour/binding/simple
http://psychedelicnekopunch.com/archives/2013
https://stackoverflow.com/questions/41503758/conversion-of-time-duration-type-microseconds-value-to-milliseconds
https://all.undo.jp/asr/1st/document/10_04.html


