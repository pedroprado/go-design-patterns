
# 1.Creational Design Patterns

## 1.1.Builder

#### O que é 

* Design pattern que ajuda na hora de criar objetos **complexos**

* Construção de objetos **step-by-step**

* Oferece uma API para que este passo à passo seja entedível e simples

## 1.2.Factories

#### O que é 

* Também Ajuda na criação de objetos **complexos**

* Entidade que auxilia na construção de objetos **completos** (não step-by-step como o Builder)

#### Exemplos

* Factory Function: uma função auxiliar, que ajuda na criação do objeto

* Structura Factory: um objeto (struct) que toma conta da criação de outro objeto

## 1.3.Prototype

#### O que é

* Um objeto ou método que ajuda a copiar/clonar objetos já **completamente inicializados**

* Isto resolve o problema que pode existir quando algums objetos copiados não lidam com o problema de ponteiros (os objetos original e copiado apontam para um mesmo ponteiro em determinado campo).

* Um prototype, portanto, que ajuda a **copiar e customizar** um objeto

#### Exemplos

* Deep Copy: copia-se campo a campo toda vez que for necessário 

* Copy Method: método ou função que copia um objeto

* Copy Serialization: copia através de um objeto serializado

## 1.4.Singleton

#### O que é

* Usado para componentes que **só precisam existir um** no sistema e sua construção é muito custosa

* A ideia é dar a mesma referência do componente para todos os outros componentes que o utilizam

* O objetivo é **previnir a inicialização de novas instâncias desnecessárias**

#### Exemplos

* Factories

* Repositories

#### Problemas

* Quebra o Dependency Inversion Principle: quando usamos o Singleton, este é uma implementação concreta. Logo, os componentes que o usam não dependem de uma interface polimórfica, o que quebra o DIP. Isso também dificulta a criação de testes unitários, já que a implementação não pode ser mockada.

* Superando o problema: o problema de DIP com Singleton pode ser superado com o uso de interfaces polimórficas (sendo o singleton a única instância concreta possível desta interface)


# 2.Structural Design Patterns

## 2.1.Adapter

#### O que é 

* Um componente que transforma uma interface A em uma interface desejada B

#### Exemplo

* Temos uma interface de **VectorImage** que representa uma imagem composta por vários pontos. Queremos imprimir os pontos que compõe esta imagem através da função **DrawImage()**. Porém, esta função não aceita um VectorImage, mas apenas uma **interface RasterImage**. Logo, precisamos de um Adaptador, ou seja, um elemento que é desta inferface RasterImage (implementa-a) e que consegue extrair as informações necessárias de um VectorImage para imprimi-lo com DrawImage()

## 2.2.Bridge

#### O que é

* Resolve o problema de **Produto Cartesiano**, desacoplando a abstração da implementação

* Se temos dois **tipos**, cada um com 3 **subtipos**, sem a utilização do Bridge, teríamos um total de 2x3 = 6 classes para resolver o problema

* Essa situação causa o Produto Cartesiano pois tenta-se entender os tipos como dimensões diferentes. Com o padrão Bridge, **uma dessas dimensões é eliminada e usada como dependência** na outra (através da **composição**)

* Para situações simples (como a aprensentada), isso não muda muito o número de classes que serão necessárias. Mas para números maiores, o uso do padrão faz toda diferença: vamos supor o caso de **10 tipos x 10 subtipos** = **100** classes sem o uso do padrão Bridge e **20** classes com o seu uso

#### Exemplo

* Dois tipos de formas: Círculo e Retângulo

* Dois tipos de renderização: Raster e Vector

* Se não usássemos o Bridge (injenção de dependência de interfaces polimórifcas), seria preciso 4 classes para satisfazer o problema de **diversos tipos de renderização para diferentes formas**

* Com o Bridge, criamos uma interface polimórfica (Renderer) e injetamos às formas, desacoplando da implementação

* Por isso, o Bridge pode ser pensado como uma **forma mais forte de encapsulamento**

## 2.3.Composite

#### O que é

* Resolve o problema de uniformização de comportamento para valores e coleções

* Esta uniformização significa **ter a mesma interface**

#### Exemplo

 * Temos uma classe **MinhaClasse**. Queremos que tanto um único objeto de MinhaClasse como uma coleção dela (List<MinhaClasse>), ao usar um método **MeuMétodo**, tenahm o mesmo comportamento.

## 2.4.Decorator

#### O que é
* Adiciona comportamento à objetos e classes sem a necessidade de modificá-los (o que quebraria o OCP)

* Para isso, gera-se uma NOVA interface, adicinando funcionalidades e mantendo uma referência a interface antiga

* Duas formas de fazer: usando Composição ou Herança (em Go não existe Herança). 

## 2.5.Facade

#### O que é

* O objetivo deste padrão é prover uma **interface simples** e esconder a complexidade do software

* Balanceamento entre usa complexidade e usabilidade/apresentação

* Esconde um sistema complexo e o expões por uma API simples

#### Exemplo

## 2.6.Flyweight

#### O que é

* Padrão que ajuda a evitar a redundância no armazenamento de dados

* Ajuda a otimizar a quantidade de memória utilizada

* A ideia é guardar **dados comuns** em algum componente externo e acessá-los via índice ou ponteiro

#### Exemplo

* Queremos guardar textos formatados por capitalização. Inicialmente, precisamos de uma lista de boolean para guardar a informação de qual letra deve ser capitalizada.

* Com o uso do Flyweight evitamos o uso desta lista, que pode ser muito grande para textos grandes.

## 2.7.Proxy

#### O que é

* Uma classe que funciona como uma interface para um recurso em particular, não permitindo o acesso direto à ele.

* O objetivo é mudar o comportamento de uma interface sem mudá-la (preserva a interface).

* Adiciona camadas de validação a comportamentos já existentes. 

* Geralmente usa herança.

* Proxy Vs Decorator: Proxy usa a mesma interface x Decorator provê uma nova interface "melhorada"
 
#### Exemplo 

* Problema: temos duas classes, Car e um Driver. Car, o qual tem o comportamento "dirigir". Não queremos ferir o OCP (logo não podemos modificar nem Driver nem Car). Mas queremos impedir que o Driver dirija se tiver idade < 16.  

* Solução: Adicionar um CarProxy, que tem o mesmo comportamento de Car (implementa a mesma interface), mas adicionando a verificação da idade no método Drive(). 


# 3.Behavioral Design Patterns

## 3.1.Chain of Responsability

#### O que é

* Cadeia de processamento de um comando
 
* Cada componente da cadeia processa o evento de sua maneira

* Deve ter um processamento padrão e ter a habilidade de ser terminado

#### Exemplo: Comand Query Separation

* Command: pede para uma ação ocorrer

* Query: pede por informação

* Exemplo: Broker Chain


## 3.2.Command

#### O que é

* É um objeto que representa uma instrução: Encapsula os detalhes de uma operação em um objeto separado

* Pode ser usado para manter o histórico/rastreamento de ações que ocorreram sobre determinado objeto

* Comandos podem ser desfeitos (apenas os bem sucedidos deveriam poder)

#### Exemplo

* Conta do Banco: este objeto tem um método Depositar. Para manter o rastreamento de todas as operações que ocorreram sobre a conta usamos um Commando que faz a chamada do método Depositar.


## 3.3.Interpreter

* Toda entrada de texto precisa ser colocada em estruturas de dados

* Texto é transformado em ***Tokens Léxicos*** (Lexing) e depos ***Processado*** (Parsing)

## 3.4.Iterator


#### O que é

* Objeto que facilita a travessia de uma estrutura de dados em particular (como por exemplo uma ***Tree Traversal*** é um iterator de uma binary tree)

* Mantém uma referência ao ***elemento atual***, e deve saber como se mover para o ***próximo elemento***

#### Exemplo

* Em Go a iteração pode ser facilmente implementada usando o ***for range*** em ***slices*** ou um simples ***for*** sobre elementos de ***canais***

* Objetos que facilitam a travessia de estrutura de dados complexas, como Binary Trees (Tree Traversal)

## 3.5.Mediator

#### O que é

* Componente que facilita a comunicação entre outros componentes

* Os componentes se comunicando não precisam saber da existência uns dos outros (isto é, ter referência um do outro)

* Todos os componentes referenciam o Mediator. O Mediator mantém a referência de todos os componentes

#### Exemplo

* Pessoas se comunicando em Chat Room (este é o Mediator): Pessoas precisam se comunicar mas ***não se comunicam diretamente***. Elas se comunicam com uma sala (Chat Room) e esta comunica as pessoas. Assim, as pessoas precisam ter a referência do ***Chat Room*** apenas. E Chat Room precisa ter a referência de todas as pessoas.

## 3.6.Memento

#### O que é 

* Token que representa um estado do sistema (como um snapshot)

* Permite o ***retorno ao estado prévio***

* Usado pelo ***Cliente***

* Mementos não pode mudar estado (read only) nem possui comportamento (métodos)

* Pode ser usado como uma maneira de guardar histórico das mudanças (o que pode ser muito pesado para sistemas grandes)

#### Exemplo

* Conta Bancária: uma conta bancária tem a possibilidade de fazer um depósito, aumentando seu saldo. Queremos guardar os valores do saldo da conta bancária a cada depósito, de maneira que seja possível *retornar* a este saldo a qualquer momento. Com um *Memento*, guardamos o *estado da conta bancária a cada depósito*.


## 3.7.Observer

#### O que é

* Uma maneira de alguém ser informado quando uma mudança ocorre

* Envolve dois componentes: 
    * observer: objeto que quer **ser informado** da mudança (subscriber/client)
    * observable: objeto que gera os eventos de notificação para **todos** os observers (aceita subscriptions)

* Evento: dados com tipo dinâmico (para poderem ser processados por todos os observers, na sua própria maneira)

#### Exemplo

* Pessoal que fica doente: quando uma pessoa fica doente ela vai ao médico, o qual passa a tomar conta da pessoa. Nesse sentido, a ***pessoa é o observable*** (notifica o médico quando sente alguma coisa) e o ***médico o observer*** (fica de olho no paciente)

## 3.8.State

#### O que é

* Objetos que mudam o comportamento dependendo do seu estado

* Algum tipo de evento é responsável por mudar o estado do objeto

* State Machine: objeto responsável por gerenciar a mudança de estado de objetos

#### Exemplo

* Temos uma Lâmpada que pode ter o estado Ligada/Desligada. Quando está Ligada, ela só pode ser Desligada, e quando está Desligada só pode ser Ligada.

* Para isso pode-se definir uma interface de estado, que mude o comportamento da lâmpada de acordo com a implementação da interface que a mesma usa 

## 3.9.Strategy

#### O que é

* Ajuda a separar os conceitos de **Alto Nível** (parte específica) e **Baixo Nível** (partes comuns)

* A estratégia deve ser **injetada** no algoritmo de alto nível

* Usa **composição**

#### Exemplo

* Impressão: Queremos imprimir um texto, o qual pode ser em Markdown ou Html. Para isso, definimos o algoritmo de alto nível (a Impressão em si) no qual deve ser injetado uma interface de ***Estratégia de Impressão***, a qual pode ser uma das duas mencionadas.

## 3.10.Template Method

#### O que é 

* Similar ao padrão **Strategy** (este pode usar Herança em outras linguagens, o que não é possível em Go, já que o mesmo não tem herança)

* Pode ser feito com **Herança** (o que não existe em Go) ou com uma simples **function**

* No caso da função, tudo o que esta precisa fazer é depender de e usar uma interface

#### Exemplo

* Jogo de Xadrez: Queremos simular um jogo básico de Xadrez, o qual é iniciado por uma função **PlayGame** (este é o Template Method). Esta função deve depender e usar de uma interface para executar o jogo.

## 3.11.Visitor

#### O que é

* Usado quando queremos aplicar uma nova funcionalidade a um tipo e seus subtipos (todos os objetos daquele tipo e sua hieraquia). Não queremos violar o **OCP** nem o **SRP**

* Para adicionar essa nova funcionalidade usa-se um elemento, o **Visitor**

* Tipos de Visitor: 
    * Intrusive: viola o **OCP**. **Modifica** interfaces já existentes
    * Reflective: **função** que recebe o tipo desejado e adiciona a funcionalidade. Faz **checagem de tipo** 
    * Classic: usa single ou double **dispatch**
        * single dispatch => Decidir que método ou função chamar baseado no tipo de objeto **sobre o qual** a função ou método vai ser chamada (1 parâmetro).
        * double dispatch => Decidir qual método ou função chamar baseado no tipo do objeto **que chama** e tipo de objeto **sobre o qual** a função vai ser chamada (2 parâmetros)

#### Example

* Vamos supor que temos uma Inteface que representa uma expressão matemática **Expression**

* Nós já criamos várias implementações concretas, como as expressões de Adição e Multiplicação

* Agora precisamos criar uma maneira de Imprimir **todas as Expressões** (tipos e subtipos) de maneira particular

* Pode-se fazer isso:
    * Usando a maneira intrusiva: modificando a interface Expression, adicionando o método Print
    * Usando a maneira reflexiva: adicionando uma **função** que recebem as expressões, que faz **checagem dos tipos** e adiciona a funcionalidade de imprimir (de acordo com o tipo)
    * Usando a maneira clássica: faz-se o objeto (Expressão) aceitar um **Visitante**, implementando o método **Accept**. Os Visitantes são os responsáveis por determianarem o que vai ser feito sobre o visitado (respeitando o SRP)