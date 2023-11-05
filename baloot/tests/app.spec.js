// @ts-check
const { test, expect } = require('@playwright/test');

test('game play', async ({ page }) => {
  await page.goto('http://localhost:3000');

  const passe   = await page.getByRole('button', { name: 'Passe' });
  const trefle  = await page.getByRole('button', { name: 'Trefle' });
  const carreau = await page.getByRole('button', { name: 'Carreau' });
  const coeur   = await page.getByRole('button', { name: 'Coeur' });
  const pique   = await page.getByRole('button', { name: 'Pique' });
  const cent    = await page.getByRole('button', { name: 'Cent' });
  const tout    = await page.getByRole('button', { name: 'Tout' });

  await expect(passe).toHaveText(/Passe/);
  await expect(trefle).toHaveText(/Trefle/);
  await expect(carreau).toHaveText(/Carreau/);
  await expect(coeur).toHaveText(/Coeur/);
  await expect(pique).toHaveText(/Pique/);
  await expect(cent).toHaveText(/Cent/);
  await expect(tout).toHaveText(/Tout/);

  const takingCards = await page.locator('#takingCards > div');
  await expect(takingCards).toHaveCount(5)

  await tout.click()
  const playingCards = await page.locator('#playingCards > div');
  await expect(playingCards).toHaveCount(8)
});
