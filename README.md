# GolangNumberCoreCPU.Example
Golang core CPU parallel processing example usage

На powershell запустите Docker command: <br/>
docker  run  --rm -v ${PWD}/:/home/src/ -it  golang bash <br/>
cd /home/src/ <br/>
cat urls.txt | go run main.go <br/>

Примерно результат так выглядит: <br/>

root@0df6dfed1e77:/home/src# cat urls.txt | go run main.go <br/>
https://habr.com/ru/post/215117/;200;350868;593.41ms <br/>
https://ru.wikipedia.org/wiki/HTTP;200;262550;701.93ms <br/>
https://developer.mozilla.org/ru/docs/Web/HTTP;200;582333;613.00ms <br/>
https://ru.bmstu.wiki/HTTP_(Hypertext_Transfer_Protocol);200;74610;830.29ms <br/>
https://proselyte.net/tutorials/http-tutorial/introduction/;200;22971;726.83ms <br/>
http://pki.gov.kz/index.php/ru/ncalayer;200;41434;253.40ms <br/>
https://www.speedcheck.org/ru/wiki/http/;200;134778;666.90ms <br/>
http://adilet.zan.kz/rus;200;65046;120.28ms <br/>
https://wiki.merionet.ru/servernye-resheniya/3/protocol-http/;200;35610;278.83ms <br/>
https://www.opennet.ru/docs/RUS/http/;200;11506;330.94ms <br/>
<порядковый номер паралельного потока запросов>:<число запросов> <br/>
0:6 <br/>
1:4 <br/>
