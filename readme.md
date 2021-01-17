

## 2. Composite

#### O que é

* Resolve o problema de uniformização de comportamento para valores e coleções

#### Exemplo

 * Temos uma classe MinhaClasse. Queremos que tanto um único objeto de MinhaClasse como uma coleção dela (List<MinhaClasse>), ao usar um método MeuMétodo, tenahm o mesmo comportamento.

## 2. Decorator

#### O que é
* Adiciona comportamento à objetos e classes, sem a necessidade de modificá-los (o que quebraria o OCP)

* Para isso, gera-se uma NOVA interface, adicinando funcionalidades e mantendo uma referência a interface antiga

* Duas formas de fazer: usando Composição ou Herança. 


#### Exemplo

* Temos duas classes, Bird e Lizard. Queremos uma terceira classe (Dragon) que pode fazer tudo o que Bird e Lizard fazem, e mais algumas coisas 


## 3. Proxy Design Pattern

#### O que é

* Uma classe que funciona como uma interface para um recurso em particular, não permitindo o acesso direto à ele.

* O objetivo é mudar o comportamento de uma interface SEM mudá-la (preserva a interface).

* Adiciona camadas de validação a comportamentos já existentes. 

* Geralmente usa herança.

 
#### Exemplo 

* Problema: temos duas classes, Car e um Driver. Car, o qual tem o comportamento "dirigir". Não queremos ferir o OCP (logo não podemos modificar nem Driver nem Car). Mas queremos impedir que o Driver dirija se tiver idade < 16.  

* Solução: Adicionar um CarProxy, que herda o comportamento de Car, adicionando a verificação da idade no método Drive(). 