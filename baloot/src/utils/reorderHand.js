const reorderCards = (dragItem, dragOverItem, cards, setCards) => {
  const index = Array.from(dragItem.current.parentNode.children).indexOf(
    dragItem.current,
  );
  const index2 = Array.from(dragOverItem.current.parentNode.children).indexOf(
    dragOverItem.current,
  );
  const copyListItems = [...cards];
  const dragItemCount = copyListItems[index];
  copyListItems.splice(index, 1);
  copyListItems.splice(index2, 0, dragItemCount);
  dragItem.current = null;
  dragOverItem.current = null;
  setCards(copyListItems);
};

export default reorderCards;
