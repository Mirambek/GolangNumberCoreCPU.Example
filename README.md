# GolangNumberCoreCPU.Example
Golang core CPU parallel processing example usage

На powershell запустите Docker command: 
docker  run  --rm -v ${PWD}/:/home/src/ -it  golang bash
cd /home/src/
cat urls.txt | go run main.go

Примерно результат так выглядит:

root@0df6dfed1e77:/home/src# cat urls.txt | go run main.go
https://habr.com/ru/post/215117/;200;350868;593.41ms
https://ru.wikipedia.org/wiki/HTTP;200;262550;701.93ms
https://developer.mozilla.org/ru/docs/Web/HTTP;200;582333;613.00ms
https://ru.bmstu.wiki/HTTP_(Hypertext_Transfer_Protocol);200;74610;830.29ms
https://proselyte.net/tutorials/http-tutorial/introduction/;200;22971;726.83ms
http://pki.gov.kz/index.php/ru/ncalayer;200;41434;253.40ms
https://www.speedcheck.org/ru/wiki/http/;200;134778;666.90ms
http://adilet.zan.kz/rus;200;65046;120.28ms
https://wiki.merionet.ru/servernye-resheniya/3/protocol-http/;200;35610;278.83ms
https://www.opennet.ru/docs/RUS/http/;200;11506;330.94ms
<порядковый номер паралельного потока запросов>:<число запросов>
0:6
1:4
