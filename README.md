get intermidiate cert file
> extract from any web browser that run on any joined domain machine
   by select intermediate CA

generate pem file
~~~
openssl x509 -inform DES -in dev2-intermidiate.cer -out im.pem -text
~~~

run command
> go run main.go -im=im.pem