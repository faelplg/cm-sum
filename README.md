Curso de Microsserviços:
# Desenvolvimento de Aplicações Modernas e Escaláveis com Microsserviços

## Módulo: DEVOPS

### Exercício: Desafio de CI
* Autor: Rafael Goulart
  * E-mail cadastrado: rafaelgoulart@residuall.org
  * E-mail pessoal: faelplg@gmail.com
* Nome do artefato: `cm-sum`
* [Repositório do GitHub (este)](https://github.com/faelplg/cm-sum)

#### Tarefa 1
* Foi desenvolvida uma aplicação simples que chama uma função `sum` enviando
 dois argumentos e essa função retorna a soma de ambos;
* A função já envia como padrão os argumentos 5 e 5 conforme enunciado;

### Tarefa 2
* Foi criado um teste unitário que testa se o resultado da soma é 10;

### Tarefa 3
* O processo de CI foi criado utilizando o Google Cloud Build, cujo gatilho é
acionado quando um pull request é feito a este repositório por alguém com permissão;
* Os passos do CI envolvem:
    1. Verificar a versão do GO;
    2. Executar o arquivo de testes;
    3. Fazer o build do programa em Go;
    4. Fazer o build da imagem em Docker;
    5. Enviar imagem para o Container Registry;
