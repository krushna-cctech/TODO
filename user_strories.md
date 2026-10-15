# User Stories: TODO Kanban Board

## Goal

As an end user, I want to organize my TODOs on a Kanban board so that I can see what needs attention, what I am working on, and what I have finished.

## Stories

### 1. View my tasks on a board

As an end user, I want to see my tasks arranged in To Do, In Progress, and Done columns so that I can understand my workload at a glance.

Acceptance criteria:
- Each task appears in the column that matches its current status.
- The board displays the task name on each card.
- When a column has no tasks, it is still visible and clearly shown as empty.

### 2. Create a task in a workflow stage

As an end user, I want to create a task and choose its starting column so that I can place it where it belongs in my workflow.

Acceptance criteria:
- I can enter a task name and select To Do, In Progress, or Done.
- If I do not select a status, the task is added to To Do.
- The new task appears in the selected column after it is saved.
- I receive a clear message if the task cannot be saved.

### 3. Move a task between columns

As an end user, I want to move a task to another column so that the board reflects its latest progress.

Acceptance criteria:
- I can move a task from any column to any other column, including back to an earlier stage.
- Moving a task updates its status and displays it in the destination column.
- The updated status remains in place when I reload or revisit the board.
- If the status update fails, I am told and the board continues to show the task's saved status.

### 4. Update task details

As an end user, I want to edit a task's name so that the board stays accurate when my plans change.

Acceptance criteria:
- I can update a task name from its card or task details.
- The updated name is shown on the board after saving.
- A blank task name is not accepted, and I am told how to correct it.

### 5. View task details

As an end user, I want to open a task and see its details so that I can confirm what the work involves before acting on it.

Acceptance criteria:
- I can open a task from its board card.
- The task name and current status are displayed.
- If a description or due date is available for the task, it is displayed as well.

### 6. Delete a task

As an end user, I want to delete a task I no longer need so that it does not clutter my board.

Acceptance criteria:
- I can choose to delete a task from its card or task details.
- I am asked to confirm before the task is deleted.
- After confirmation, the task is removed from the board.
- If deletion fails, I am notified and the task remains visible.

### 7. Find tasks across the board

As an end user, I want to search for a task by name so that I can find it quickly when my board grows.

Acceptance criteria:
- Search matches task names without requiring an exact full-name match.
- Matching tasks remain in their respective status columns.
- I can clear the search to see all tasks again.
- A clear empty-result state is shown when no tasks match.

### 8. Continue using my board after returning

As an end user, I want my tasks and their statuses to be saved so that I can continue where I left off when I return.

Acceptance criteria:
- Created, updated, moved, and deleted tasks are reflected in the saved board data.
- Reloading the board shows the latest saved tasks and statuses.
- If tasks cannot be loaded, I see an error state and can retry.


