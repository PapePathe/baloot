import PlayingCardsView from './../baloot/components/PlayingCardsView';

export default {
  title: 'Baloot/PlayingCardsView',
  component: PlayingCardsView,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export const CinqCouleurs = {
  args: {
    cards: [
      {Couleur: "Pique", Genre: "V"},
      {Couleur: "Pique", Genre: "9"},
      {Couleur: "Pique", Genre: "A"},
      {Couleur: "Pique", Genre: "10"},
      {Couleur: "Pique", Genre: "7"},
    ]
  }
}
