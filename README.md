Напишите код, реализующий пайплайн, работающий с целыми числами и состоящий из следующих стадий:

    1. Стадия фильтрации отрицательных чисел (не пропускать отрицательные числа).
    2. Стадия фильтрации чисел, не кратных 3 (не пропускать такие числа), исключая также и 0.
    3. Стадия буферизации данных в кольцевом буфере с интерфейсом, соответствующим тому, который был дан в качестве задания в 19 модуле. В этой стадии предусмотреть опустошение буфера (и соответственно, передачу этих данных, если они есть, дальше) с определённым интервалом во времени. Значения размера буфера и этого интервала времени сделать настраиваемыми (как мы делали: через константы или глобальные переменные).

Написать источник данных для конвейера. Непосредственным источником данных должна быть консоль.

Также написать код потребителя данных конвейера. Данные от конвейера можно направить снова в консоль построчно, сопроводив их каким-нибудь поясняющим текстом, например: «Получены данные …».

При написании источника данных подумайте о фильтрации нечисловых данных, которые можно ввести через консоль. Как и где их фильтровать, решайте сами.