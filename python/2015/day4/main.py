import hashlib

def hash_starts_with(number: int, digest: str, starts_with) -> int:
    string_to_hash = digest + str(number)
    hashed = hashlib.md5(str(string_to_hash).encode('utf-8')).hexdigest()
    return hashed.startswith(starts_with)


def mine(data: str, target: str) -> int:
    number = 0
    while not hash_starts_with(number, data, target):
        number += 1
    return number


def main():
    input = 'bgvyzdsv'
    print(f'Part 1: {mine(input, "00000")}')
    print(f'Part 2: {mine(input, "000000")}')


if __name__ == "__main__":
    main()
