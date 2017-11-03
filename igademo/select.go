package ga

/////// Elete Select /////////////////

func EleteSelection(p []Gene, num int) []Gene {
    pop := CopyPop(p)
    BubbleSort(pop)
    return pop[len(p)-num:]
}
