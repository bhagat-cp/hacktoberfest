def countWaysToJump(arr, n):
 
    count_jump = [0 for i in range(n)]
 
    for i in range(n - 2, -1, -1):
     
        if (arr[i] >= n - i - 1):
            count_jump[i] += 1
 
        j = i + 1
        while(j < n-1 and j <= arr[i] + i):
 
            if (count_jump[j] != -1):
                count_jump[i] += count_jump[j]
            j += 1

        if (count_jump[i] == 0):
            count_jump[i] = -1
     
 
    for i in range(n):
        print(count_jump[i], end = " ")

arr = [1, 3, 5, 8, 9, 1, 0, 7, 6, 8, 9]
n = len(arr)
countWaysToJump(arr, n)
