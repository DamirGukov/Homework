package main

import "fmt"

type Titles struct {
	EarlyYears        string
	DuringTheCivilWar string
	WritersWork       string
	LastYears         string
}

func (t Titles) earlyYears() string {

	return t.EarlyYears
}

func (t Titles) civilWar() string {
	return t.DuringTheCivilWar

}

func (t Titles) writersWork() string {
	return t.WritersWork
}

func (t Titles) lastYears() string {
	return t.LastYears
}

func main() {
	t := Titles{
		EarlyYears: "\nThe future classic of American literature was born in 1835 in the village of Florida (Missouri). " +
			"The father died when the boy was 13 years old, the mother lived a long life and died at the age of 87. " +
			"In addition to Sam, the family had 3 more children: two boys and a girl. After the death of his father, Sam's older brother Orion became the head of the family." +
			" It was he who opened the family business: he began to publish a newspaper. " +
			"Samuel also worked in the publishing house, first as a typesetter, and then as a journalist." +
			" As a journalist, he traveled the country, visited St. Louis and New York." +
			"After working for his brother for a while, Samuel realized that the river was \"calling\" him. " +
			"He became a pilot on a steamship. He liked the work, but the Civil War led to the disappearance of the private shipping company. Samuel was forced to start looking for a livelihood again." +
			"It is known that at the very beginning of the Civil War, the future writer became a member of the Masonic lodge, although he always treated the brotherhood with humor.\n",
		DuringTheCivilWar: "\nFor a time, Samuel fought in the ranks of the people's militia, but after his brother was made secretary to the governor of Nevada, he left with him to the West." +
			"In Nevada, Sam worked at a mine as a silver miner. Then he got a job at the Territorial Enterprise newspaper." +
			"In 1864, Sam moved to San Francisco, where he began working for several newspapers at once.\n",
		WritersWork: "\nAfter 1870, Twain came to grips with writing. " +
			"Also at this time, he began teaching at a number of universities in the United States and England. " +
			"Twain was an excellent speaker and his lectures were incredibly popular. " +
			"In his later works, the author spoke out against racism and imperialism, criticized current US senators, and spoke negatively about presidents." +
			" By the way, his novel The Adventures of Huckleberry Finn was banned several times, as they believed that the words and expressions used by the authors were unliterary, and many scenes were too naturalistic.\n",
		LastYears: "\nIn recent years, the writer's financial affairs have been greatly shaken, but the situation was saved by the oil tycoon Henry Rogers, who became a close friend of the writer." +
			"Mark Twain greatly influenced the character of the American businessman and made him a real philanthropist and philanthropist." +
			"Roger, at the request of the writer, organized several charitable foundations that sponsored educational programs for African Americans and for children with disabilities." +
			"The writer was buried several times. After another obituary, Mark Twain even uttered the catchphrase that the rumors of his death were greatly exaggerated." +
			"He died in 1910 from an attack of angina pectoris. It is known that he was born in the year when Halley's comet passed over the earth, he also “left” with it, since in 1910 it again passed by the Earth (by the way, the writer actually predicted his death).\n",
	}

	fmt.Println(t.earlyYears())
	fmt.Println(t.civilWar())
	fmt.Println(t.writersWork())
	fmt.Println(t.lastYears())

}
