# li
```
$ # li should be useful after classic "sort | uniq -c" combo
$ go get github.com/reinerRubin/li/cmd/li
$ perl -le 'map { print } split "", "bddudddumddduuuddddm"' | sort | uniq -c | li
 12 ■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■  d
  5 ■■■■■■■■■■■■                    u
  2 ■■■■■                           m
  1 ■■                              b
```
