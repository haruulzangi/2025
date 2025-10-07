import base64
import pickle
from random import randint, choice
import socketserver
import signal

WELCOME = """
*************** Монгол Баатрын Зэвсгийн Дэлгүүрт Тавтай Морилно Уу! **********
*                                        *
*  Та олон төрлийн монгол зэвсэг худалдаж авч хүчээ нэмэгдүүлээрэй...        *
*  зарим зэвсэг хямд бөгөөд муу чанартай байж болно!                         *
*                                        *
******************************************************************************
"""

MONGOL_WEAPONS = [
  "Нум",       
  "Сум",        
  "Илд",        
  "Сэлэм",      
  "Балдан",     
  "Буу",        
  "Бахь",       
  "Дайны_туг",  
  "Хуяг",       
  "Дуулга"      
]

MONGOL_HERO_SKILLS = [
  {
    "name": "Баямбадалай",
    "skill": "Тэнгэрийн Сум",
    "role": "Алсын харваач",
    "description": "Тэнгэрийн ивээлтэй мэт холын зайнд онож харвах чадвартай. Алсад байгаа дайснуудыг сандаргаж, арын эгнээг цөөлнө.",
    "health": 90,
    "balance": 100,
    "power": 15
  },
  {
    "name": "Ялгуун",
    "skill": "Тэнгэрийн дуудлага",
    "role": "Урам зориг хайрлагч",
    "description": "Дайчдад урам зориг хайрлаж, хүчийг нь нэмэгдүүлдэг. Эсрэг талаа айдаст автуулж, довтолгооны эрчийг сулруулна.",
    "health": 85,
    "balance": 120,
    "power": 12
  },
  {
    "name": "Тэнгис",
    "skill": "Галт Морин Дайралт",
    "role": "Морин цэрэг",
    "description": "Гал мэт хурдтай морьтойгоороо дайсны эгнээг сөрөн орж, хүчтэй доргилтоор дайснаа тарааж хаяна.",
    "health": 110,
    "balance": 90,
    "power": 18
  },
  {
    "name": "Энхцэрэн",
    "skill": "Чонон Сүүдэр",
    "role": "Тагнуул",
    "description": "Хээр талын чонын адил сүүдэр мэт нуугдаж, дайсны хөдөлгөөнийг мэдэж ирсэн мэдээллээр тулааныг чиглүүлнэ.",
    "health": 100,
    "balance": 80,
    "power": 99
  },
  {
    "name": "Ууганбаяр",
    "skill": "Хар илд",
    "role": "Ойрын тулалдааны мастер",
    "description": "Илдээ салхи мэт эргэлдүүлж, нэгэн зэрэг хэд хэдэн дайсныг шархдуулдаг.",
    "health": 95,
    "balance": 70,
    "power": 20
  },
  {
    "name": "Баттулга",
    "skill": "Хуягны Сүнс",
    "role": "Хамгаалагч",
    "description": "Бүх довтолгооныг өөр дээрээ авч, нөхдөө хамгаалдаг. Хуяг нь дайсны сум, илдийг тэсвэрлэнэ.",
    "health": 140,
    "balance": 60,
    "power": 8
  },
  {
    "name": "Эрдэнэ-Оч",
    "skill": "Цуст Нум",
    "role": "Анчин",
    "description": "Онож харвах болгонд нь дайсны хүчийг сорж, өөрийн хүчийг нэмэгдүүлдэг.",
    "health": 80,
    "balance": 85,
    "power": 22
  },
  {
    "name": "Тэмүүжин",
    "skill": "Эрийн Дором",
    "role": "Гардан тулалдагч",
    "description": "Богино зайд тулалдахад гаргууд. Дором хэмээх модон бамбай, богино сэлмээрээ дайсныг хашиж устгана.",
    "health": 105,
    "balance": 75,
    "power": 16
  },
  {
    "name": "Дэлгэрмөрөн",
    "skill": "Талын Шуурга",
    "role": "Тактикч",
    "description": "Тал нутгийн салхи, шуургыг бэлгэдэн ашиглаж, дайсны эгнээг бут цохихоор зохион байгуулалттай довтолгоо удирдана.",
    "health": 95,
    "balance": 110,
    "power": 13
  },
  {
    "name": "Төгөлдөр",
    "skill": "Сүлдний Хүч",
    "role": "Домогт баатар",
    "description": "Өвөг дээдсийн сүлдний ивээлээр хүч чадал, эр зоригийг хэд дахин нэмэгдүүлж, бүх цэргийг хамтад нь урамшуулна.",
    "health": 120,
    "balance": 100,
    "power": 25
  },
  {
    "name": "Дарханбаяр",
    "skill": "Талын Аянга",
    "role": "Довтолгооны мастер",
    "description": "Аянга шиг гэнэт гарч ирээд хүчтэй цохилт өгдөг. Дайсанд бэлтгэх зав өгдөггүй.",
    "health": 100,
    "balance": 95,
    "power": 17
  },
  {
    "name": "Эрхэмбаяр",
    "skill": "Илдний Сүлд",
    "role": "Сэлэмний мастер",
    "description": "Хүчирхэг сэлмээрээ дайсныг нэгэн зэрэг хэд хэдэн удаа цохиж, тэдний сүнсийг айдаст автуулдаг.",
    "health": 90,
    "balance": 105,
    "power": 19
  },
  {
    "name": "Мөнхбаатар",  
    "skill": "Галт сум",
    "role": "Харваач",
    "description": "Харвасан сум нь гал мэт хурдан дайсныг өчиггүй устгадаг.",
    "health": 130,
    "balance": 65,
    "power": 10
  },
  {
    "name": "Цогточир",  
    "skill": "Тамын Сүх", 
    "role": "Дайсныг бут цохигч",
    "description": "Тамаас ивээлтэй мэт хүчирхэг сүхээрээ дайсныг бут цохиж, тэдний сүнсийг айдаст автуулдаг.",
    "health": 115,
    "balance": 85,
    "power": 14
  },
  {
    "name": "Баярхүү",
    "skill": "Тэнгэрийн Хуяг",
    "role": "Хамгаалагч",
    "description": "Хуяг нь ямар ч дайсны довтолгоог тэсвэрлэж, нөхдөө хамгаалдаг. Хуяг нь дайсны сум, илдийг тэсвэрлэнэ.",
    "health": 150,
    "balance": 50,
    "power": 7
  }
]


class Handler(socketserver.BaseRequestHandler):

  def handle(self):
    signal.alarm(0)
    main(self.request)


class ReusableTCPServer(socketserver.ForkingMixIn, socketserver.TCPServer):
  pass


def sendMessage(s, msg):
  s.send(msg.encode('utf-8'))


def receiveMessage(s, msg):
  sendMessage(s, msg)
  return s.recv(4096).decode('utf-8').strip()


def choose_hero(s):
  sendMessage(s, "\n=============== БААТАР СОНГОХ ===============\n")
  sendMessage(s, "Дараах монголын домогт баатруудаас нэгийг сонгоно уу:\n\n")
  
  for i, hero in enumerate(MONGOL_HERO_SKILLS, 1):
    sendMessage(s, f"[{i}] {hero['name']} - {hero['skill']} ({hero['role']})\n")
    sendMessage(s, f"     {hero['description']}\n")
    sendMessage(s, f"     Эрүүл мэнд: {hero['health']}, Мөнгө: {hero['balance']}, Хүч: {hero['power']}\n\n")
  
  while True:
    try:
      choice_num = int(receiveMessage(s, "Баатрын дугаарыг сонгоно уу (1-{}): ".format(len(MONGOL_HERO_SKILLS))))
      if 1 <= choice_num <= len(MONGOL_HERO_SKILLS):
        selected_hero = MONGOL_HERO_SKILLS[choice_num - 1]
        sendMessage(s, f"\n✓ {selected_hero['name']} ({selected_hero['skill']}) -г сонголоо!\n")
        return selected_hero
      else:
        sendMessage(s, f"\nБуруу сонголт! 1-ээс {len(MONGOL_HERO_SKILLS)} хүртэл дугаар оруулна уу.\n")
    except ValueError:
      sendMessage(s, "\nТоо оруулна уу!\n")
    except Exception as e:
      sendMessage(s, f"\nАлдаа гарлаа: {e}\n")


class Weapon():

  def __init__(self, name, price, power):
    self.name = name
    self.price = price
    self.power = power 


class WeaponShop():

  def __init__(self, s):
    self.options = "[1] Зэвсгийн жагсаалт харах.\n[2] Зэвсэг худалдаж авах.\n[3] Юу ч хийхгүй.\n"
    self.storage = [
      Weapon(choice(MONGOL_WEAPONS) + f"_{i}", randint(10, 50), randint(5, 15))
      for i in range(8)
    ]
    self.s = s

  def listItems(self):
    sendMessage(self.s, "\nЗэвсгийн нэр    | Үнэ  | Хүч\n")
    sendMessage(self.s, "-----------------------------------\n")
    for weapon in self.storage:
      sendMessage(
        self.s,
        f"{weapon.name:<18} | {weapon.price:<4} | {weapon.power}\n")

  def parseStorageByName(self, name):
    for index in range(len(self.storage)):
      if (self.storage[index].name == name):
        return index
    return None

  def sellTo(self, name, warrior):
    index = self.parseStorageByName(name)

    if index != None:
      weapon = self.storage[index]
      price = weapon.price
      if (warrior.balance >= price):
        warrior.balance -= price
        warrior.weapons.append(weapon)
        self.storage.remove(weapon)
        sendMessage(self.s, f"\n{weapon.name} амжилттай худалдаж авлаа!\n")
      else:
        sendMessage(self.s, "\nТаны мөнгө хүрэхгүй байна\n")
    else:
      sendMessage(self.s, "\nТийм зэвсэг олдсонгүй\n")


class MongolWarrior():

  def __init__(self, s, hero_data):
    self.options = "\n[1] Зэвсэг ашиглах.\n[2] Зэвсгээ хаях.\n[3] Юу ч хийхгүй.\n"
    self.name = hero_data['name']
    self.skill = hero_data['skill']
    self.role = hero_data['role']
    self.description = hero_data['description']
    self.health = hero_data['health']
    self.balance = hero_data['balance']
    self.power = hero_data['power']
    self.weapons = []
    self.s = s

  def use_weapon(self):
    if len(self.weapons) == 0:
      sendMessage(self.s, "\nТанд ашиглах зэвсэг алга!\n")
      return
      
    weapon = self.weapons[0]

    if weapon.price > 25:
      self.power += weapon.power
      self.health += 5
      sendMessage(self.s, f"\n{weapon.name} маш сайн зэвсэг байлаа! Хүч нэмэгдлээ: +{weapon.power}\n")
    else:
      self.health -= 15
      sendMessage(self.s, f"\n{weapon.name} муу зэвсэг байна. Цус хасагдлаа!\n")

    self.weapons.remove(weapon)

    if self.health <= 0:
      sendMessage(self.s, "Та ялагдлаа! Тоглоом дуусав!\n")
      exit(1)

  def free(self):
    self.weapons = []
    return "Зэвсгүүдыг хаялаа!"

  def get_status(self):
    return f"\nБаатар: {self.name} ({self.skill}) | Эрүүл мэнд: {self.health} | Мөнгө: {self.balance} | Хүч: {self.power} | Зэвсгийн тоо: {len(self.weapons)}\n"


def main(s):
  shop = WeaponShop(s)
  
  sendMessage(s, WELCOME)
  
  selected_hero = choose_hero(s)
  warrior = MongolWarrior(s, selected_hero)
  
  sendMessage(s, warrior.get_status())

  while True:
    sendMessage(s, "\nСайн байна уу, юу хийхийг хүсэж байна?\n")
    sendMessage(s, shop.options)

    try:
      option = int(receiveMessage(s, "\n> "))

      if option == 1:
        shop.listItems()
      elif option == 2:
        sendMessage(s, "\nХудалдаж авах зэвсэг?\n")
        weapon_name = receiveMessage(s, "\n> ")
        shop.sellTo(weapon_name, warrior)
      else:
        sendMessage(s, "")

      sendMessage(s, warrior.get_status())

      if len(warrior.weapons) > 0:
        sendMessage(s, "\nҮйлдэл?\n")
        sendMessage(s, warrior.options)
        option = int(receiveMessage(s, "\n> "))

        if option == 1:
          warrior.use_weapon()
        elif option == 2:
          data = receiveMessage(s, '\nТулаанд үргэлж хамт байсан зэвсэгтээ дайчны ёсоор сүүлийн үгээ хэлнэ үү: ')
          data = base64.b64decode(data)
          _warrior = pickle.loads(data)
          sendMessage(s, f'\nҮр дүн: {_warrior.free()}\n')
        else:
          pass
      else:
        sendMessage(s, "\nТанд ашиглах зэвсэг байхгүй байна.\n")

    except KeyboardInterrupt:
      sendMessage(s, "\n\nГарч байна...")
      exit(1)
    except Exception as e:
      sendMessage(s, f"\nДата боловсруулахад алдаа гарлаа: {e}\n")


if __name__ == '__main__':
  socketserver.TCPServer.allow_reuse_address = True
  server = ReusableTCPServer(("0.0.0.0", 11337), Handler)
  server.serve_forever()