# buchladen

---

- [DESIGN OVERVIEW](#design-overview)

## DESIGN OVERVIEW

This is an experimental design of a microservice-oriented software, the
main goal of which is to interact with users and keep track of items
they sell.

We need to have an API to store various users of our service as well as
manipulate such data:

```
                    _____
+---------+        /     \
|  Users  | <---> ( MySQL )
|   API   |       (\_____/)
+---------+       (_______)
```

[OAuth](https://www.rfc-editor.org/rfc/rfc6749) authorisation framework
shall be used for user authentication. Since we’ve taken the
microservice-oriented approach, we shall create a new NoSQL
database—[Apache Cassandra](https://cassandra.apache.org) in this
case—to be used in conjunction with our OAuth API to store access tokens
and whatnot:

```
  _______
 /       \
( Cassan- )
(   dra   )
(\_______/)
(_________)
(_________)

     ^
     |
     |
     v
+---------+
|  OAuth  |
|   API   |
+---------+
```
