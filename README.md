# ProjetoInterativoV

Repositorio unico destinado ao desenvolvimento e manutencao dos microservicos do Projeto iniciado no PI V, 1o semestre do ano de 2023 no curso de Engenharia da Computacao do SENAC.

## Autores

* Felipe Fernandes Mandelli
* Gregorio Alves
* Guilherme
* Murillo Tiberio Costelini
* Thiago Andrade

## Para Rodar o código

Primeiro, é necessario alterar o arquivo `cmd\gateway\config\config.yaml`. No campo Adress, altere `localhost`para o seu ip. Devendo ficar dessa forma:

``` yaml
  ADDRESS: "127.0.0.1:9015"

```

Certifique-se de que a porta 9015 da sua maquina esteja habilitada a receber protocolo http pelo Windows Firewall.

Iniciar > Windows Defender Firewall > Configuração avançada > Regras de Entrada > Nova Regra

Esta será a porta em que o ESP ira fazer as requisições http para realizar a chamada de cada aula e passar a tag lida no cartão.
