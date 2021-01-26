

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
