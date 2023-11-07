const { test, expect } = require('@playwright/test');

test('game take carreau', async ({ page }) => {
  await page.goto('http://localhost:3000');

  const takingCards = await page.locator('#takingCards > div');
  await expect(takingCards).toHaveCount(5)

  const playingCards = await page.locator('#playingCards > div');
  await expect(playingCards).toHaveCount(0)

  const carreau = await page.getByRole('button', { name: 'Carreau' });
  await carreau.click()

  await expect(takingCards).toHaveCount(5)
});

test('game take coeur', async ({ page }) => {
  await page.goto('http://localhost:3000');

  const takingCards = await page.locator('#takingCards > div');
  await expect(takingCards).toHaveCount(5)

  const playingCards = await page.locator('#playingCards > div');
  await expect(playingCards).toHaveCount(0)

  const carreau = await page.getByRole('button', { name: 'Coeur' });
  await carreau.click()

  await expect(takingCards).toHaveCount(5)
});



test('game take tout', async ({ page }) => {
  await page.goto('http://localhost:3000');
  page.on('console', msg => {
    if (msg.type() === 'error') {
      console.error(`Error text: "${msg.text()}"`);
    } else {
      console.log(msg)
    }
  });

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
