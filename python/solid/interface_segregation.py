import abc
from abc import abstractmethod


class Operation(abc.ABCMeta):
    @abstractmethod
    def print(cls):
        pass

    @abstractmethod
    def scan(cls):
        pass

    @abstractmethod
    def fax(cls):
        pass


class MultiFunctionMachine(Operation):  # all of these methods are okay because this obj can handle them
    def print(cls):
        pass

    def scan(cls):
        pass

    def fax(cls):
        pass


class OldMachine(Operation):  # this will break the ISP rule
    def print(cls):
        pass

    def scan(cls):
        pass  # this is not ok because cant handle this operation

    def fax(cls):
        pass  # this is not ok because cant handle this operation

# should separate the interfaces for solution
