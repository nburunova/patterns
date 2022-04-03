# Abstract method

```mermaid
classDiagram
    class Member
    <<interface>> Member
    Member : Name()
    Member : Health()
    Member : Type()

    Member <|-- Pirate
    Pirate : Name()
    Pirate : Health()
    Pirate : Type()

    Member <|-- Knight
    Knight : Name()
    Knight : Health()
    Knight : Type()

    class Creator
    <<interface>> Creator
    Creator : CreateSquad()

    Creator <|-- PiratesGang
    PiratesGang : CreateSquad()
    Creator <|-- KnightsPlatoon
    KnightsPlatoon : CreateSquad()
```
