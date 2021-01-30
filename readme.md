

## 1. Builder

## 2. Factories

## 3. Prototype

## 4. Singleton

## 5. Adapter

## 6. Bridge

## 7. Composite

#### O que é

* Resolve o problema de uniformização de comportamento para valores e coleções

#### Exemplo

 * Temos uma classe MinhaClasse. Queremos que tanto um único objeto de MinhaClasse como uma coleção dela (List<MinhaClasse>), ao usar um método MeuMétodo, tenahm o mesmo comportamento.

## 8. Decorator

#### O que é
* Adiciona comportamento à objetos e classes, sem a necessidade de modificá-los (o que quebraria o OCP)

* Para isso, gera-se uma NOVA interface, adicinando funcionalidades e mantendo uma referência a interface antiga

* Duas formas de fazer: usando Composição ou Herança. 


## 9. Facade

## 10.Flyweight



#### Exemplo

* Temos duas classes, Bird e Lizard. Queremos uma terceira classe (Dragon) que pode fazer tudo o que Bird e Lizard fazem, e mais algumas coisas 


## 11. Proxy

#### O que é

* Uma classe que funciona como uma interface para um recurso em particular, não permitindo o acesso direto à ele.

* O objetivo é mudar o comportamento de uma interface SEM mudá-la (preserva a interface).

* Adiciona camadas de validação a comportamentos já existentes. 

* Geralmente usa herança.
 
#### Exemplo 

* Problema: temos duas classes, Car e um Driver. Car, o qual tem o comportamento "dirigir". Não queremos ferir o OCP (logo não podemos modificar nem Driver nem Car). Mas queremos impedir que o Driver dirija se tiver idade < 16.  

* Solução: Adicionar um CarProxy, que herda o comportamento de Car, adicionando a verificação da idade no método Drive(). 


## 12. Chain of Responsability

#### O que é

* Cadeia de processamento de um comando
 
* Cada componente da cadeia processa o evento de sua maneira

* Deve ter um processamento padrão e ter a habilidade de ser terminado

#### Comand Query Separation

* Command: pede para uma ação ocorrer

* Query: pede por informação

* Exemplo: Broker Chain


## 13. Command

#### O que é

* Representação de uma ação: X deve alterar Z (X manda um comando à Z)

* É um objeto que representa uma instrução: Encapsuça os detalhes de uma operação em um objeto separado

* Usado para manter o histórico/rastreamento de ações que ocorreram sobre determinado objeto

* Comandos devem poder serem desfeitos

* Apenas comandos bem sucedidos devem poder serem desfeitos

#### Exemplo

* Conta do Banco: este objeto tem um método Depositar. Para manter o rastreamento de todas as operações que ocorreram sobre a conta usamos um Commando que faz a chamada do método Depositar.


## 14. Interpreter

* Toda entrada de texto precisa ser colocada em estruturas de dados

* Texto é transformado em ***Tokens Léxicos*** (Lexing) e depos ***Processado*** (Parsing)

## 15. Iterator


#### O que é

* Objeto que facilita a travessia de uma estrutura de dados em particular (como por exemplo uma ***Tree Traversal*** é um iterator de uma binary tree)

* Mantém uma referência ao ***elemento atual***, e deve saber como se mover para o ***próximo elemento***

#### Exemplo

* Em Go a iteração pode ser facilmente implementada usando o ***for range*** em ***slices*** ou um simples ***for*** sobre elementos de ***canais***

* Objetos que facilitam a travessia de estrutura de dados complexas, como Binary Trees (Tree Traversal)

## 16. Mediator

#### O que é

* Componente que facilita a comunicação entre outros componentes

* Os componentes se comunicando não precisam saber da existência uns dos outros (ter referência)

* Todos os componentes referenciam o Mediator. O Mediator mantém a referência de todos os componentes

#### Exemplo

* Pessoas se comunicando em Chat Room (este é o Mediator): Pessoas precisam se comunicar mas ***não se comunicam diretamente***. Elas se comunicam com uma sala (Chat Room) e esta comunica as pessoas. Assim, as pessoas precisam ter a referência do ***Chat Room*** apenas. E Chat Room precisa ter a referência de todas as pessoas.

## 17. Memento

#### O que é 

* Token que representa um estado do sistema (como um snapshot)

* Permite o ***retorno ao estado prévio***

* Usado pelo ***Cliente***

* Mementos não pode mudar estado (read only) nem possui comportamento (métodos)

* Pode ser usado como uma maneira de guardar histórico das mudanças (o que pode ser muito pesado para sistemas grandes)

#### Exemplo

* Conta Bancária: uma conta bancária tem a possibilidade de fazer um depósito, aumentando seu saldo. Queremos guardar os valores do saldo da conta bancária a cada depósito, de maneira que seja possível *retornar* a este saldo a qualquer momento. Com um *Memento*, guardamos o *estado da conta bancária a cada depósito*.
