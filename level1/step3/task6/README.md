Секретный язык
Ограничение времени	1 секунда
Ограничение памяти	64 Мб
Ввод	стандартный ввод
Вывод	стандартный вывод
Программист Арсений придумал секретный язык: он убирает из слов все буквы а и о. Поговаривают, что ещё ни одна секретная служба не смогла расшифровать его тайнопись.

Напишите программу, которая будет получать на вход слово и выводить его шифр на секретном языке.

Формат ввода
Строка

Формат вывода
Строка

Пример
Ввод	Вывод
подарок
пдрк
Примечания
При проверке символов а и о нужно использовать такую конструкцию: letter == 'а'. С помощью одинарных кавычек в Go задаются литералы руны (подробнее об этом вы узнаете на следующих занятиях).