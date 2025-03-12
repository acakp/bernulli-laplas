# Решение задач на вероятность с помощью формул Бернулли и Муавра-Лапласа

Пусть проводится _n_ независимых испытаний, при каждом испытании может произойти одно и _k_ событий, т.е. эти события образуют полную группу.

  Ω = {A<sub>1</sub>, A<sub>2</sub>, ..., A<sub>k</sub>}
  
При этом возникновение каждого события в любом из _n_ испытаний имеет постоянную вероятность

  P(A<sub>1</sub>) = P<sub>1</sub>
  P(A<sub>2</sub>) = P<sub>2</sub>
  ...
  P(A<sub>k</sub>) = P<sub>k</sub>
  
Тогда вероятность того, что событие A<sub>1</sub> возникнет n<sub>1</sub> раз, событие A<sub>2</sub> возникнет n<sub>1</sub> раз, ... A<sub>k</sub> возникнет n<sub>k</sub> раз определятеся по формуле Бернулли:

![P=\frac{n!}{n_{1}! \cdot n_{2}! \cdot ... \cdot n_{k}!} \cdot P_1^{n_1} \cdot P_2^{n_2} \cdot ... \cdot P_k^{n_k}](https://latex.codecogs.com/svg.image?\bg{white}P=\frac{n!}{n_{1}!\cdot&space;n_{2}!\cdot...\cdot&space;n_{k}!}\cdot&space;P_1^{n_1}\cdot&space;P_2^{n_2}\cdot...\cdot&space;P_k^{n_k})

При этом _n_ = _n_<sub>1</sub> + _n_<sub>2</sub> + ... + _n_<sub>k</sub>. Данная формула следует из правила вычисления перестановок с повторениями.

Однако, часто требуется сосчитать вероятности событий, в которых количество испытаний велико, поэтому потребуется сосчитать большие факториалы и степени. Для этого существуют приближенные формулы.

Формулы Муавра-Лапласа:

![P=\frac{\phi(x)}{\sqrt{npq}} \qquad x=\frac{k-nP}{\sqrt{npq}} \qquad \phi(x)=\frac{1}{\sqrt{2\pi}}e ^{-\frac{x^2}{2}}](https://latex.codecogs.com/svg.image?\bg{white}P=\frac{\phi(x)}{\sqrt{npq}}\qquad&space;x=\frac{k-nP}{\sqrt{npq}}\qquad\phi(x)=\frac{1}{\sqrt{2\pi}}e^{-\frac{x^2}{2}})


## Пример

Студенту на 2 курсе предстоит сдать 8 экзаменов. Вероятность сдать на "5" - 10% (P<sub>1</sub> = 0.1), вероятность сдать на "4" - 20% (P<sub>2</sub> = 0.2), вероятность сдать на "3" или "2" - 70% (P<sub>3</sub> = 0.7).
Требуется найти вероятность того, что будет две пятерки, две четверки и четыре двойки или тройки.

Для решения этой задачи в программу будут введены следующие значения:

```
Введите количество экспериментов (n): 8
Введите количество вероятностей (k): 3
Введите количество событий для каждой вероятности:
n1: 2
n2: 2
n3: 4
Введите вероятность для каждого события (от 0 до 1):
P1: 0.1
P2: 0.2
P3: 0.7
```

Результат будет следующий:
```
Результат по формуле Бернулли: 0.04033680
Результат по формуле Муавра-Лапласа: 0.17296146

Решение по формуле Бернулли:
(8!/(2!*2!*4!))*0.10^2*0.20^2*0.70^4 = 0.04033680

Решение по формуле Муавра-Лапласа:
x = (2 - 8*0.10)/sqrt(8*0.10*(1-0.90)) = 1.41
φ = (1/sqrt(2*π))*exp^(-(1.41^2/2)) = 0.15
Р = 0.15/sqrt(8*0.10*(0.90)) = 0.17296146
```
