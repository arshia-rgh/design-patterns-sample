from threading import Lock, Thread


class SingletonMeta(type):
    _instances = {}
    _lock = Lock()

    def __call__(cls, *args, **kwargs):
        with cls._lock:
            if cls not in cls._instances:
                instance = super().__call__(*args, **kwargs)

                cls._instances[cls] = instance

        return cls._instances[cls]


class Singleton(metaclass=SingletonMeta):
    value = None

    def __init__(self, value):
        self.value = value

    def some_func(self):
        pass

def test_singleton(value):
    single = Singleton(value)
    print(single.value)

if __name__ == "__main__":
    process_1 = Thread(target=test_singleton, args=("FOO",))
    process_2 = Thread(target=test_singleton, args=("BAR",))

    process_1.start()
    process_2.start()
