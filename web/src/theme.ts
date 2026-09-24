export {};

type Theme = 'light' | 'dark';

const storageKey = 'bpl-theme';
const toggle = document.getElementById('themeToggle') as HTMLButtonElement | null;
const label = toggle?.querySelector('.theme-label');

function getInitialTheme(): Theme {
	const storedTheme = localStorage.getItem(storageKey);
	if (storedTheme === 'light' || storedTheme === 'dark') return storedTheme;
	return window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light';
}

function applyTheme(theme: Theme) {
	document.documentElement.dataset.theme = theme;
	const nextTheme = theme === 'dark' ? 'light' : 'dark';

	if (toggle) {
		toggle.setAttribute('aria-label', `Switch to ${nextTheme} mode`);
		toggle.setAttribute('title', `Switch to ${nextTheme} mode`);
	}
	if (label) label.textContent = nextTheme === 'dark' ? 'Dark' : 'Light';
}

applyTheme(getInitialTheme());

toggle?.addEventListener('click', () => {
	const nextTheme: Theme = document.documentElement.dataset.theme === 'dark' ? 'light' : 'dark';
	localStorage.setItem(storageKey, nextTheme);
	applyTheme(nextTheme);
});
