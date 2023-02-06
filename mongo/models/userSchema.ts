const mongoose = require("mongoose");

/**
 *  Mongo schema to define type: User.
 *  A document (user) ID is automatically generated
 */
const UserSchema = mongoose.Schema({
  email: {
    type: String,
    required: true,
  },

  displayName: {
    type: String,
    required: true,
  },

  // Channels the user has joined
  channels: {
    type: Array,
    required: false,
    default: [],
  },
});

module.exports = mongoose.model("User", UserSchema);
