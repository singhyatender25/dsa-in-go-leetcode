func maxArea(height []int) int {
    
    i:=0
    j:=len(height)-1

    maxArea:= 0

    for (i<j) {
        w:= j-i
        h:= min(height[i],height[j])
        area:= w*h
        maxArea = max(maxArea,area)
        
        if height[i] > height[j] {
            j--
        } else {
            i++
        }
    }
    return maxArea
}