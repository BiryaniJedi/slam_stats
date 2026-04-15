export const sanitizeName = (name: string): string => {
	return name.trim().replaceAll(' ', '+');
};
