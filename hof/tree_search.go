package main

import "log"

type State int

type States []State

// GoalP is a predicate that returns true if the state is a goal state.
type GoalP func(State) bool

// SuccessorsP returns a list of successor states for a given state.
type SuccessorsP func(State) States

// Combiner combines two lists of states into one.
// It determines the order(search strategy) in which the states are visited.
type Combiner func(States, States) States

func treeSearch(states States, goalP GoalP, succ SuccessorsP, combiner Combiner) State {
	log.Println("states:", states)
	if len(states) == 0 {
		return -1
	}

	var first State
	first, states = states[0], states[1:]
	if goalP(first) {
		return first
	}
	return treeSearch(combiner(succ(first), states), goalP, succ, combiner)
}

func bfsTreeSearch(start State, goalP GoalP, succ SuccessorsP) State {
	return treeSearch(States{start}, goalP, succ, prependOthers)
}

func dfsTreeSearch(start State, goalP GoalP, succ SuccessorsP) State {
	return treeSearch(States{start}, goalP, succ, appendOthers)
}

// appendOthers is a combiner that appends others to succ.
func appendOthers(succ States, others States) States {
	return append(succ, others...)
}

// prependOthers is a combiner that prepends others to succ.
func prependOthers(succ States, others States) States {
	return append(others, succ...)
}

// stateIsGoal returns a GoalP that returns true if the state is equal to the goal.
func stateIsGoal(goal State) GoalP {
	return func(s State) bool {
		return s == goal
	}
}

// binaryTree returns the two children of a state.
func binaryTree(s State) States {
	return States{s * 2, s*2 + 1}
}

// finiteBinaryTree returns a SuccessorsP that returns the two children of a state
func finiteBinaryTree(n State) SuccessorsP {
	return func(s State) States {
		return filter(
			binaryTree(s),
			func(item State) bool {
				return item <= n
			},
		)
	}
}

// filter returns a new list of elments that satisfy the predicate.
func filter[T any](s []T, pred func(T) bool) []T {
	var result []T
	for _, e := range s {
		if pred(e) {
			result = append(result, e)
		}
	}
	return result
}

func main() {
	start := State(1)
	goal := State(10)
	succ := finiteBinaryTree(goal)
	log.Println("bfsTreeSearch:", bfsTreeSearch(start, stateIsGoal(goal), succ))
}
