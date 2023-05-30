package common

func Covert_Slice_To_Map(sl []string) map[string]struct{} {
	// []string -> map[string]struct{}
	fmt_map := make(map[string]struct{}, len(sl))
	for _, tag := range sl {
		fmt_map[tag] = struct{}{}
	}
	return fmt_map
}

func Check_key_exists(fmt_map map[string]struct{}, tag string) bool {
	// check map key exists
	_, elelement_res := fmt_map[tag]
	return elelement_res
}

func CheckIndexIsExceedListLen(podIndex int, podList []string) bool {
	if len(podList) >= podIndex+1 {
		return true
	} else {
		return false
	}
}
