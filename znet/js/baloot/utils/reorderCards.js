const reorderCards = (
  dragItem,
  dragOverItem,
  playingCards,
  setPlayingCards,
) => {
  const index = Array.from(dragItem.current.parentNode.children).indexOf(
    dragItem.current,
  );
  const index2 = Array.from(dragOverItem.current.parentNode.children).indexOf(
    dragOverItem.current,
  );
  const copyListItems = [...playingCards];
  const dragItemCount = copyListItems[index];
  copyListItems.splice(index, 1);
  copyListItems.splice(index2, 0, dragItemCount);
  dragItem.current = null;
  dragOverItem.current = null;
  setPlayingCards(copyListItems);
};

export default reorderCards;
