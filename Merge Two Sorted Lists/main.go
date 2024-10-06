/*
You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.

Решение:
-созадать узел
-слияние двух списков
-цикл for
-добавление оставшихся узлов

сложность алгоритма составляет O(n + m), где n и m — длины двух списков
*/
package main

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// Создаем новый "виртуальный" узел, который будет использоваться как начало нового списка
	dumm := &ListNode{}
	current := dumm

	// Слияние двух списков
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	// Добавляем оставшиеся узлы
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	// Возвращаем следующий узел от dumm, который является началом объединенного списка
	return dumm.Next
}
