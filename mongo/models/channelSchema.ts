import * as mongoose from "mongoose";

/**
 *  Mongo schema to define type: Room.
 *  A document (channel) ID is automatically generated
 */
const ChannelSchema = new mongoose.Schema({
  name: {
    type: String,
    required: false,
    default: "New Channel",
  },
  private: {
    type: Boolean,
    required: false,
    default: true,
  },
  users: {
    type: Array,
    required: true,
    default: [],
  },
});

export default mongoose.model("Channel", ChannelSchema);
