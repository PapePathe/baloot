import  CardView from './../baloot/components/CardView';

export default {
  title: 'Baloot/CardView',
  component: CardView,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
};

export const ValetCarreau = {
  args: { Couleur: "Carreau", Genre: "V", width: 32, height: 64}
}

export const ValetPique= {
  args: { Couleur: "Pique", Genre: "V", width: 24, height: 44}
}

export const DixPique = {
  args: { Couleur: "Pique", Genre: "10"}
}


export const DameCoeur = {
  args: { Couleur: "Coeur", Genre: "10"}
}
