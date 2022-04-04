# Observer

## ConcreteObserver may have a ConcreteObservable instance and call ConcreteObservable.GetState() inside Update()

```mermaid
classDiagram
    class Observable
    <<interface>> Observable
    Observable : Add(Observable o)
    Observable : Remove(Observable o)
    Observable : Notify()

    Observable <|-- ConcreteObservable
    ConcreteObservable : Add(Observable o)
    ConcreteObservable : Remove(Observable o)
    ConcreteObservable : Notify()
    ConcreteObservable : GetState()

    class Observer
    <<interface>> Observer
    Observer : Update()

    Observer <|-- ConcreteObserver1
    ConcreteObserver1 : -concreteObservable instance
    ConcreteObserver1 : Update()
    Observer <|-- ConcreteObserver2
    ConcreteObserver2 : -concreteObservable instance
    ConcreteObserver2 : Update()

    ConcreteObserver1 ..> ConcreteObservable
    ConcreteObserver2 ..> ConcreteObservable
```

## ConcreteObserver receive updated state as an argument for Update(args)

```mermaid
classDiagram
    class Observable
    <<interface>> Observable
    Observable : Add(Observable o)
    Observable : Remove(Observable o)
    Observable : Notify()

    Observable <|-- ConcreteObservable
    ConcreteObservable : Add(Observable o)
    ConcreteObservable : Remove(Observable o)
    ConcreteObservable : Notify()

    class Observer
    <<interface>> Observer
    Observer : Update()

    Observer <|-- ConcreteObserver1
    ConcreteObserver1 : Update(newState)
    Observer <|-- ConcreteObserver2
    ConcreteObserver2 : Update(newState)
```