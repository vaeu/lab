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
