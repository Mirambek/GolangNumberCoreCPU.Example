# GolangNumberCoreCPU.Example
Golang core CPU parallel processing example usage

На powershell запустите Docker command: <br/>
docker  run  --rm -v ${PWD}/:/home/src/ -it  golang bash <br/>
cd /home/src/ <br/>
cat urls.txt | go run main.go <br/>

Примерно результат так выглядит: <br/>
root@3ae9a49b828c:/home/src# cat urls.txt | go run main.go<br/>
https://ru.wikipedia.org/wiki/HTTP;200;262550;936.09ms<br/>
https://habr.com/ru/post/215117/;200;350763;1363.06ms<br/>
https://developer.mozilla.org/ru/docs/Web/HTTP;200;582333;999.69ms<br/>
https://ru.bmstu.wiki/HTTP_(Hypertext_Transfer_Protocol);200;74610;1104.83ms<br/>
https://proselyte.net/tutorials/http-tutorial/introduction/;200;22971;792.50ms<br/>
http://pki.gov.kz/index.php/ru/ncalayer;200;41430;209.17ms<br/>
http://www.edu.gov.kz/;200;533;486.80ms<br/>
https://www.speedcheck.org/ru/wiki/http/;200;130758;1277.38ms<br/>
https://wiki.merionet.ru/servernye-resheniya/3/protocol-http/;200;35458;337.94ms<br/>
https://www.opennet.ru/docs/RUS/http/;200;11506;304.30ms<br/>
https://www.w3.org/Protocols/HTTP/1.1/rfc2616bis/draft-lafon-rfc2616bis-03.html;200;745969;1795.10ms<br/>
http://adilet.zan.kz/rus;200;65046;3238.78ms<br/>
https://flaviocopes.com/http/;200;66365;680.99ms<br/>
https://www.extrahop.com/resources/protocols/http/;200;80118;2342.66ms<br/>
<порядковый номер паралельного потока запросов>:<число запросов><br/>
0:7<br/>
1:7<br/>