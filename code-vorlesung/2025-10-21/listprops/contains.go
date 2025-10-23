package listprops

// Contains liefert true, falls die Liste l
// den String x enthält.
func Contains(l []string, x string) bool {
<<<<<<< HEAD
	// TODO
=======
>>>>>>> 027ee69e8af7d7b335bee7cc01448d4dbdecd9df
	for i := 0; i < len(l); i++ {
		if l[i] == x {
			return true
		}
	}
	return false
}
