# easy crypto challenge by zjzoloo 💀
# 💀 maybe skipping would be better 💀

from binascii import hexlify as _H
from gmpy2 import *
import math as _M, os as _O, sys as _S

if _S.version_info < (3, 9):
    _M.gcd, _M.lcm = gcd, lcm

💀 = False

🦄 = mpz(_H(open('flag.txt').read().strip().encode()), 16)
🌱 = mpz(_H(_O.urandom(32)).decode(), 16)
🎲 = random_state(🌱)

def 🧩(🎲, 🪙):
    return next_prime(mpz_urandomb(🎲, 🪙) | (1 << (🪙 - 1)))

def 🧨(🎲, 🪙, 🐍=16):
    🪐 = mpz(2)
    🦖 = [🪐]
    while 🪐.bit_length() < 🪙 - 2 * 🐍:
        f = 🧩(🎲, 🐍)
        🦖.append(f)
        🪐 *= f

    🔢 = (🪙 - 🪐.bit_length()) // 2

    while True:
        a, b = 🧩(🎲, 🔢), 🧩(🎲, 🔢)
        tmp = 🪐 * a * b
        if tmp.bit_length() < 🪙: 🔢 += 1; continue
        if tmp.bit_length() > 🪙: 🔢 -= 1; continue
        if is_prime(tmp + 1):
            🦖 += [a, b]; 🪐 = tmp + 1; break
    🦖.sort()
    return (🪐, 🦖)

🔑 = 0x10001

while True:
    p, pf = 🧨(🎲, 1024, 16)
    if len(pf) != len(set(pf)): continue
    q, qf = 🧨(🎲, 1024, 17)
    if len(qf) != len(set(qf)): continue
    if 🔑 not in (pf + qf): break

n = p * q
m = _M.lcm(p - 1, q - 1)
d = pow(🔑, -1, m)
c = pow(🦄, 🔑, n)

print(f'n = {n.digits(16)}')
print(f'c = {c.digits(16)}')

# p factor : a9b8ac82034aca5e36d082411b6d02fb9ae2abd2c6a0761e601ce6686ccd221695516e7a8b548885a92c97ccb2297bebdd5cafb42e5ca105593fae843b9f298f0a9939f4d90ce6ed903bb68948c3f479fdf47f4dfb6b07ff11fd09becc32ebaa15da0b4753b0569c0c6a35d852944db30273b374afbf0d75483aa191b5135e661d183fefd2b550af6debb2da54f33e48bcefebbefc6be55d02b0e0859a81b0904fffb41472257185add81e141580e78eba0074a69e3048607d444711399ab62da163191fd57bc560355f81f67a7cee281e08b36bea19a6ab7cf79930aee63c2672d461971720271dbcf3c8fda2a7fc9431787948e3527d568307762abbc7b951
# q factor : 2023ef922965d23c5b75389a6307ba372883be7f5718d4e2401b21fd6e476e17ed7d1f773f95f2909a9f80ef7d1e7b325f647706d380d2196d24be6990e05b14bf97549269f6a709aba7fb71538b3a5ccb51ec526c16835d2b02e670b1ec9856e216b2b2cf52bc136c74ccafd994c0282efe5f2bc4dbca7b3323724190284410bab5f0e726eb226b79a675491a560fb6d4bb36cc051e524f6215ac52dadd7d97d0f7671c333d5665b148380165017b03c67bc9046d699097462614f7470afa3aeffada1ffb356d13261fcc0b9ee36a4bf9fedfe006413516dd1c0db53c7d35818fa334104df77fc390b6e9a6985c7c47f45c2245fd78c22ecc593bb02a3d5e03
