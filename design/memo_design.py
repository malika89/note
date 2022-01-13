# coding:utf-8

"""
备忘录模式，在不破坏封闭的前提下，捕获一个对象的内部状态，并在该对象之外保存这个状态。这样以后就可将该对象恢复到原先保存的状态

应用场景：
  草稿箱
  游戏进度保存
  事务
"""

import random


class Memento:
    vitality = 0
    attack = 0
    defense = 0

    def __init__(self, vitality, attack, defense):
        self.vitality = vitality
        self.attack = attack
        self.defense = defense
        
        
class GameCharacter(object):
    vitality = 0
    attack = 0
    defense = 0
    
    def display_state(self, status="游戏中"):
        print(f"{status}中状态如下：")
        print('生命值:%d' % self.vitality)
        print('攻击值:%d' % self.attack)
        print('防御值:%d' % self.defense)
        print("\n")
        
    def init_state(self, vitality, attack, defense):
        self.vitality = vitality
        self.attack = attack
        self.defense = defense
        
    def save_state(self):
        return Memento(self.vitality, self.attack, self.defense)
    
    def recover_state(self, memento):
        self.vitality = memento.vitality
        self.attack = memento.attack
        self.defense = memento.defense
        
        
class FightCharactor(GameCharacter):
    def fight(self):
        self.vitality -= random.randint(1,10)
        self.attack += random.randint(2,6)


if __name__ == "__main__":
    game_chrctr = FightCharactor()
    # 打斗角色初始化
    game_chrctr.init_state(100, 79, 60)
    game_chrctr.display_state("开局")
    memento = game_chrctr.save_state()
    
    # 打斗游戏中
    game_chrctr.fight()
    game_chrctr.display_state()
    
    # 异常恢复
    game_chrctr.recover_state(memento)
    game_chrctr.display_state("上一次")