#!/usr/bin/env python3
from Cryptodome.Util.number import bytes_to_long, getPrime
from random import randint, seed
import json

flag = b""

def encrypt(msg, nbit):
    m = bytes_to_long(msg)
    p = getPrime(nbit)
    assert m < p

    e = 65537
    t = 5  

    C = [randint(0, p - 1) for _ in range(t - 1)] + [pow(m, e, p)]

    def poly_val(x):
        res = 0
        for i in range(t):
            res += C[i] * pow(x, t - i - 1, p)
        return res % p

    seed(42)
    PT = [(i+1, poly_val(i+1)) for i in range(t)]
    
    return e, p, PT

if __name__ == "__main__":
    nbit = 512
    enc = encrypt(flag, nbit)

    # Save output to file as JSON
    with open("output.txt", "w") as f:
        json.dump(enc, f)

    print("Encryption output saved to output.txt")
