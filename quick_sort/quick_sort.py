def swap(lst, index1, index2):
    lst[index1], lst[index2] = lst[index2], lst[index1]

def pivot(lst, pivot_index, end_index):
    swap_index = pivot_index

    for i in range(pivot_index + 1, end_index + 1):
        if lst[i] < lst[pivot_index]:
            swap_index += 1
            swap(lst, swap_index, i)

    swap(lst, pivot_index, swap_index)
    return swap_index
    
def quick_sort_helper(lst, left, right):
    if left < right:
        pivot_index = pivot(lst, left, right)
        quick_sort_helper(lst, left, pivot_index - 1)
        quick_sort_helper(lst, pivot_index + 1, right)
    return lst

def quick_sort(lst):
    return quick_sort_helper(lst, 0, len(lst) - 1)

print(quick_sort([1, 2, 7, 8, 3, 10, 11, 4, 9, 5, 6]))