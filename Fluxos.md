# Fluxos

## Fluxo de informações do ESP32 para o banco de dados

* O ESP32 envia uma solicitação HTTP para a API REST, informando que um aluno está presente na aula.

* A API REST recebe a solicitação HTTP e verifica se o aluno está matriculado na aula e se já não foi registrado como presente.
* Se o aluno estiver matriculado e ainda não tiver sido registrado como presente, a API REST adiciona uma entrada no banco de dados indicando que o aluno está presente na aula.

## Fluxo de informações do banco de dados para a interface do usuário

* A interface do usuário solicita ao servidor web uma lista de alunos presentes na aula.
* O servidor web envia uma solicitação à API REST para obter a lista de alunos presentes na aula.
* A API REST retorna a lista de alunos presentes na aula para o servidor web.
* O servidor web envia a lista de alunos presentes na aula para a interface do usuário.

## Fluxo de informações da interface do usuário para o banco de dados

* O professor seleciona um aluno na lista de alunos presentes na aula.
* A interface do usuário envia uma solicitação HTTP para a API REST, informando que o professor está marcando o aluno como ausente.
* A API REST recebe a solicitação HTTP e atualiza o registro do aluno no banco de dados, marcando-o como ausente.

## Fluxo de informações do banco de dados para a interface do usuário em consulta

* A interface do usuário solicita ao servidor web uma lista atualizada de alunos presentes na aula.
* O servidor web envia uma solicitação à API REST para obter a lista atualizada de alunos presentes na aula.
* A API REST retorna a lista atualizada de alunos presentes na aula para o servidor web.
* O servidor web envia a lista atualizada de alunos presentes na aula para a interface do usuário.
