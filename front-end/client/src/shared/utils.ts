
export const getDayNum = (m = 1): number => m * 24 * 60 * 60 * 1000;

export const getDayOfWeek = (orig = new Date(), i = 0) => {
	const d = new Date(+orig);
	const day = d.getDay();
	if (day === i) return d;
	d.setHours(-24 * (day - i)); 
	return d;
};

