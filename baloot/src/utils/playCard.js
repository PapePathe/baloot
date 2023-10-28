const playCard = (couleur, genre, playerID, sendMessage) => {
  if (couleur !== "" && genre !== "") {
    sendMessage(
      JSON.stringify({
        player_id: `${playerID}`,
        color: couleur,
        genre: genre,
        id: "4",
      }),
    );
  }
};

export default playCard;
